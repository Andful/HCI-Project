const el = document.getElementById('app');

document.getElementById('start').addEventListener('click', () => {
	if (screenfull.enabled) {
		screenfull.request(el);
	}
});

function start() {
	if (screenfull.enabled) {
		screenfull.request(el);
	}
}
