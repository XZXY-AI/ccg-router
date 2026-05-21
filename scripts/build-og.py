#!/usr/bin/env python3
"""Regenerate the social-preview image at docs/demo.png.

Output: 1280x640 PNG. Dark background, two terminal-like rectangles
representing Claude Code and Codex CLI both flowing into a single
center block representing ccg-router on 127.0.0.1:17180.

Run from repo root:
    python3 scripts/build-og.py
"""

from PIL import Image, ImageDraw, ImageFont
from pathlib import Path

OUT = Path(__file__).resolve().parent.parent / "docs" / "demo.png"

W, H = 1280, 640
BG = (15, 17, 22)             # near-black
PANEL = (28, 32, 40)
PANEL_EDGE = (60, 66, 78)
ACCENT = (94, 168, 255)       # cool blue
ACCENT_DIM = (54, 96, 150)
TEXT = (235, 238, 245)
TEXT_DIM = (160, 168, 180)
GREEN = (140, 220, 160)
ORANGE = (240, 170, 110)

FONT_REG = "/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf"
FONT_BOLD = "/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf"
FONT_MONO = "/usr/share/fonts/truetype/dejavu/DejaVuSansMono.ttf"
FONT_MONO_BOLD = "/usr/share/fonts/truetype/dejavu/DejaVuSansMono-Bold.ttf"


def font(path, size):
    return ImageFont.truetype(path, size)


def rounded_panel(draw, xy, fill, outline=None, width=2, radius=14):
    draw.rounded_rectangle(xy, radius=radius, fill=fill, outline=outline, width=width)


def draw_terminal(draw, xy, title, lines, accent):
    x0, y0, x1, y1 = xy
    rounded_panel(draw, xy, fill=PANEL, outline=PANEL_EDGE)
    # title bar
    bar_h = 32
    draw.rounded_rectangle((x0, y0, x1, y0 + bar_h), radius=14, fill=(40, 46, 56))
    draw.rectangle((x0, y0 + bar_h - 14, x1, y0 + bar_h), fill=(40, 46, 56))
    for i, c in enumerate([(255, 96, 92), (255, 189, 46), (39, 201, 63)]):
        cx = x0 + 16 + i * 22
        cy = y0 + bar_h // 2
        draw.ellipse((cx - 6, cy - 6, cx + 6, cy + 6), fill=c)
    draw.text((x0 + 100, y0 + 8), title, font=font(FONT_BOLD, 16), fill=TEXT_DIM)
    # body
    fnt = font(FONT_MONO, 16)
    fnt_b = font(FONT_MONO_BOLD, 16)
    ty = y0 + bar_h + 14
    for kind, t in lines:
        if kind == "prompt":
            draw.text((x0 + 16, ty), "$ ", font=fnt_b, fill=accent)
            draw.text((x0 + 36, ty), t, font=fnt, fill=TEXT)
        elif kind == "out":
            draw.text((x0 + 16, ty), t, font=fnt, fill=TEXT_DIM)
        elif kind == "ok":
            draw.text((x0 + 16, ty), t, font=fnt_b, fill=GREEN)
        ty += 22


def draw_arrow(draw, start, end, color):
    draw.line([start, end], fill=color, width=3)
    # arrow head
    ex, ey = end
    sx, sy = start
    dx = 1 if ex > sx else -1
    head = [(ex, ey), (ex - 14 * dx, ey - 7), (ex - 14 * dx, ey + 7)]
    draw.polygon(head, fill=color)


def main():
    img = Image.new("RGB", (W, H), BG)
    d = ImageDraw.Draw(img)

    # title
    d.text((48, 36), "ccg-router", font=font(FONT_BOLD, 56), fill=TEXT)
    d.text((48, 100), "One local daemon. Both Claude Code and Codex CLI talk to it.",
           font=font(FONT_REG, 22), fill=TEXT_DIM)

    # two terminals on the left, router in the middle, arrows
    term_w, term_h = 380, 190
    left_x = 48
    cc_y = 180
    cx_y = 400
    draw_terminal(d, (left_x, cc_y, left_x + term_w, cc_y + term_h),
                  "Claude Code · Anthropic-compatible",
                  [
                      ("prompt", "claude code"),
                      ("out", "ANTHROPIC_BASE_URL=127.0.0.1:17180"),
                      ("ok", "→ ccg-router"),
                  ],
                  ACCENT)
    draw_terminal(d, (left_x, cx_y, left_x + term_w, cx_y + term_h),
                  "Codex CLI · OpenAI-compatible",
                  [
                      ("prompt", "codex --resume"),
                      ("out", "OPENAI_BASE_URL=127.0.0.1:17180"),
                      ("ok", "→ ccg-router"),
                  ],
                  ORANGE)

    # router panel (center-right)
    rx0, ry0 = 580, 200
    rx1, ry1 = 900, 540
    rounded_panel(d, (rx0, ry0, rx1, ry1), fill=PANEL, outline=ACCENT, width=3, radius=18)
    d.text((rx0 + 20, ry0 + 22), "ccg-router", font=font(FONT_BOLD, 28), fill=TEXT)
    d.text((rx0 + 20, ry0 + 62), "127.0.0.1:17180", font=font(FONT_MONO, 18), fill=ACCENT)

    bullets = [
        ("/v1/messages", "Anthropic"),
        ("/v1/chat/completions", "OpenAI"),
        ("prefer-cheaper", "strategy"),
        ("local SQLite ledger", "per request"),
    ]
    by = ry0 + 110
    for path, tag in bullets:
        d.text((rx0 + 20, by), path, font=font(FONT_MONO_BOLD, 17), fill=TEXT)
        d.text((rx0 + 20, by + 22), tag, font=font(FONT_REG, 15), fill=TEXT_DIM)
        by += 56

    # arrows from terminals into router
    draw_arrow(d, (left_x + term_w, cc_y + term_h // 2), (rx0, ry0 + 120), ACCENT_DIM)
    draw_arrow(d, (left_x + term_w, cx_y + term_h // 2), (rx0, ry0 + 220), ACCENT_DIM)

    # right-side outputs (provider keys local, ledger local)
    ox0 = 950
    d.text((ox0, ry0 + 22), "stays local", font=font(FONT_BOLD, 18), fill=GREEN)
    items = [
        "~/.config/ccg-router/config.toml",
        "provider keys",
        "",
        "~/.local/share/ccg-router/ledger.db",
        "one row per request",
    ]
    iy = ry0 + 60
    for line in items:
        if line.startswith("~/"):
            d.text((ox0, iy), line, font=font(FONT_MONO, 15), fill=TEXT)
        else:
            d.text((ox0, iy), line, font=font(FONT_REG, 14), fill=TEXT_DIM)
        iy += 24

    # footer
    d.text((48, H - 50), "github.com/XZXY-AI/ccg-router", font=font(FONT_MONO_BOLD, 18), fill=ACCENT)
    d.text((48, H - 26), "Apache-2.0 · Go single binary · brew · curl · go install", font=font(FONT_REG, 14), fill=TEXT_DIM)

    OUT.parent.mkdir(parents=True, exist_ok=True)
    img.save(OUT, "PNG", optimize=True)
    print(f"wrote {OUT} ({OUT.stat().st_size} bytes)")


if __name__ == "__main__":
    main()
