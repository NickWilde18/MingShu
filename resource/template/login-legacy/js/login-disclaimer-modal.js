// Modal control logic

// 登录按钮点击处理
function handleLogin() {
    const isAccepted = localStorage.getItem('noticeAccepted') === 'true';
    if (!isAccepted) {
        showDisclaimerModal();
    } else {
        jumpToAI();
    }
}

// 显示免责声明弹窗
function showDisclaimerModal() {
    const modal = document.getElementById('disclaimer-modal');
    const content = modal.querySelector('div');

    modal.classList.remove('hidden');
    // 触发重排以应用初始状态
    void modal.offsetWidth;

    // 背景动画
    modal.classList.add('transition-opacity', 'duration-300');
    modal.classList.replace('opacity-0', 'opacity-100');

    // 内容动画
    content.classList.add('animate-modal-enter');
    content.classList.remove('opacity-0');
}

// 优化隐藏逻辑
function hideDisclaimerModal() {
    const modal = document.getElementById('disclaimer-modal');
    const content = modal.querySelector('div');

    // 内容离开动画
    content.classList.replace('animate-modal-enter', 'animate-modal-leave');

    // 背景动画
    modal.classList.replace('opacity-100', 'opacity-0');

    // 动画完成后隐藏
    setTimeout(() => {
        modal.classList.add('hidden');
        content.classList.remove('animate-modal-leave');
    }, 300);
}


// 弹窗按钮事件绑定
document.getElementById('notice-accept').addEventListener('click', jumpToAI);
document.getElementById('notice-reject').addEventListener('click', hideDisclaimerModal);

// 失焦关闭：点在遮罩但不是内容上
document.getElementById('disclaimer-modal').addEventListener('click', function (e) {
    // 只有点遮罩自身而不是内容时关闭
    if (e.target === this) {
        hideDisclaimerModal();
    }
});