// Dark/light mode logic

// 统一的过渡时间配置
const THEME_TRANSITION_DURATION = 300; // 与CSS保持一致 (0.3s = 300ms)

// 日夜切换
const themeBtn = document.getElementById('theme-toggle');

const updateThemeIcon = () => {
    const html = document.documentElement;
    const sunIcon = document.getElementById('theme-icon-sun');
    const moonIcon = document.getElementById('theme-icon-moon');
    
    if (sunIcon && moonIcon) {
        sunIcon.style.display = html.classList.contains('dark') ? 'none' : 'block';
        moonIcon.style.display = html.classList.contains('dark') ? 'block' : 'none';
    }
    
    // 夜间主卡片增加外发光
    const mainCard = document.getElementById('mainCard');
    if (mainCard) {
        if (html.classList.contains('dark')) {
            mainCard.classList.add('card-glow-dark');
        } else {
            mainCard.classList.remove('card-glow-dark');
        }
    }
}

// 平滑的主题切换函数
const smoothThemeToggle = () => {
    const html = document.documentElement;
    
    // 添加过渡效果类（如果需要）
    html.style.transition = 'all 0.3s cubic-bezier(0.4, 0.0, 0.2, 1)';
    
    // 切换主题
    if (html.classList.contains('dark')) {
        html.classList.remove('dark');
        localStorage.theme = 'light';
    } else {
        html.classList.add('dark');
        localStorage.theme = 'dark';
    }
    
    // 更新图标
    updateThemeIcon();
    
    // 在过渡完成后清理样式（可选）
    setTimeout(() => {
        // 保持过渡效果，不移除
        // html.style.transition = '';
    }, THEME_TRANSITION_DURATION);
}

themeBtn && (themeBtn.onclick = smoothThemeToggle);

// 初始化主题
(() => {
    const html = document.documentElement;
    
    // 设置初始过渡效果
    html.style.transition = 'all 0.3s cubic-bezier(0.4, 0.0, 0.2, 1)';
    
    if (localStorage.theme === 'dark' ||
        (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
        html.classList.add('dark');
    }
    updateThemeIcon();
})();

