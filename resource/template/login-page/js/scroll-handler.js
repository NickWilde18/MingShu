// Wheel/touch events


// 滚轮
window.addEventListener('wheel', function (e) {
    // 节流，避免重复
    if (page === 'main' && e.deltaY > 60) switchPage('models');
    else if (page === 'models' && e.deltaY < -85) switchPage('main');
});
// 触摸
let touchStartY = 0;
window.addEventListener('touchstart', function (e) {
    if (e.touches.length == 1)
        touchStartY = e.touches[0].clientY;
}, { passive: true });
window.addEventListener('touchend', function (e) {
    let endY = (e.changedTouches && e.changedTouches[0]) ? e.changedTouches[0].clientY : touchStartY;
    let diff = touchStartY - endY;
    if (page === 'main' && diff > 120) switchPage('models');
    else if (page === 'models' && diff < -200) switchPage('main');
    touchStartY = null;
}, { passive: false });

const observer = new IntersectionObserver(entries => {
    entries.forEach(entry => {
        entry.target.style.animationPlayState = entry.isIntersecting ? 'running' : 'paused';
    });
});
observer.observe(document.querySelector('.css-ai-bg'));