import { setLocalStorageItem, getLocalStorageItem } from './local-storage-operator.js';
import en from './locales/en.js';
import zhCN from './locales/zh-CN.js';

// 设为全局变量
window.loadLang = loadLang;

// 存储翻译内容的对象
const supportedLanguages = ['en', 'zh-CN'];
const translations = { 'en': en, 'zh-CN': zhCN };

// 动态加载翻译文件
async function loadLang() {
  // 将语言选项添加到 language-select 下拉菜单
  const languageSelect = document.getElementById('language-select');
  supportedLanguages.forEach(lang => {
    const option = document.createElement('option');
    option.value = lang;
    option.textContent = translations[lang]["lang-name"];
    languageSelect.appendChild(option);
  });

  // 设置初始语言
  languageSelect.value = loadLanguage();
  updateLanguage(languageSelect.value);
  listenLanguage();
}

function loadLanguage() {

  if (getLocalStorageItem("i18nextLng")) {
    return getLocalStorageItem("i18nextLng");
  }
  // 获取用户的首选语言
  const userLanguage = navigator.language;

  // 检查用户的首选语言是否在支持的语言列表中
  const matchedLanguage = supportedLanguages.find(lang => userLanguage.startsWith(lang));

  // 如果没有匹配的语言，则使用英语（en）
  const selectedLanguage = matchedLanguage || 'en';

  // 自动设置对应语言
  return selectedLanguage;
}

// 监听语言切换
function listenLanguage() {
  const languageSelect = document.getElementById("language-select");
  languageSelect.addEventListener("change", (e) => {
    const lang = e.target.value;
    updateLanguage(lang);
  });
}

// 更新页面语言
function updateLanguage(lang) {
  // body设置类名 en zh-CN
  document.body.classList.remove('en', 'zh-CN');
  document.body.classList.add(lang);
  setLocalStorageItem("i18nextLng", lang);
  const elements = document.querySelectorAll("[data-i18n]");
  elements.forEach((element) => {
    const key = element.getAttribute("data-i18n");
    if (translations[lang] && translations[lang][key]) {
      element.innerHTML = translations[lang][key];
    }
  });
}
