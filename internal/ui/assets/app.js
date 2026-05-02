async function load(path, id) {
  const r = await fetch(path);
  const t = await r.text();
  document.getElementById(id).textContent = t;
}

load('/api/usage/summary', 'summary-out');
load('/api/usage/window', 'window-out');
