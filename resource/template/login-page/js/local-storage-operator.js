export function setLocalStorageItem(name, value, days) {
  if (days) {
    const now = new Date();
    const expiry = now.getTime() + days * 24 * 60 * 60 * 1000; // 计算过期时间
    const item = JSON.stringify({ value, expiry }); // 包装为带过期时间的对象
    localStorage.setItem(name, item);
  } else {
    localStorage.setItem(name, value); // 直接存储纯字符串
  }
}

export function getLocalStorageItem(name) {
  const itemStr = localStorage.getItem(name);
  if (!itemStr) {
    return null; // 未找到对应的值
  }

  try {
    if (itemStr.startsWith("{")) {
      // 如果是 JSON 格式（有过期时间）
      const item = JSON.parse(itemStr);
      if (item.expiry && new Date().getTime() > item.expiry) {
        // 如果已过期，删除该项并返回 null
        localStorage.removeItem(name);
        return null;
      }
      return item.value; // 返回存储的值
    } else {
      // 如果是普通字符串，直接返回
      return itemStr;
    }
  } catch (e) {
    console.error("无法解析存储的值：", e);
    return null;
  }
}
