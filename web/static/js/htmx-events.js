document.addEventListener('htmx:configRequest', (e) => {
    e.detail.headers['Content-Type'] = 'application/json';
});

document.addEventListener('htmx:afterSwap', (e) => {
    if (e.detail.target.id === 'results') {
        const results = e.detail.target.querySelectorAll('.hit');
        results.forEach(hit => {
            hit.addEventListener('click', () => {
                console.log('Hit clicked:', hit.dataset.id);
            });
        });
    }
});

function updateMetrics(data) {
    document.getElementById('total-queries').textContent = data.queries || 0;
    document.getElementById('avg-latency').textContent = (data.latency || 0) + 'ms';
    document.getElementById('error-count').textContent = data.errors || 0;
}

setInterval(() => {
    fetch('/v1/metrics')
        .then(r => r.json())
        .then(updateMetrics)
        .catch(() => {});
}, 5000);
