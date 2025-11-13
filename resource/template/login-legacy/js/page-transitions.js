// Screen switching animations

// ---- 上下滑/滚切换，两页渐变动画，header互斥 ----
let page = 'main';
const mainScreen = document.getElementById('main-screen');
const modelsScreen = document.getElementById('models-screen');
function switchPage(to) {
    if (to === 'models' && page === 'main') {
        mainScreen.style.pointerEvents = 'none';
        mainScreen.style.opacity = 0;
        setTimeout(() => {
            mainScreen.style.display = "none";
            modelsScreen.style.display = "flex";
            setTimeout(() => {
                modelsScreen.style.opacity = 1;
                modelsScreen.style.pointerEvents = 'auto';
            }, 10);
        }, 400);
        page = 'models';
    } else if (to === 'main' && page === 'models') {
        modelsScreen.style.pointerEvents = 'none';
        modelsScreen.style.opacity = 0;
        setTimeout(() => {
            modelsScreen.style.display = "none";
            mainScreen.style.display = "flex";
            setTimeout(() => {
                mainScreen.style.opacity = 1;
                mainScreen.style.pointerEvents = 'auto';
            }, 10);
        }, 400);
        page = 'main';
    }
}
