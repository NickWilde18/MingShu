import models from './models.json';

const tbody = document.getElementById('models-table-body');
tbody.innerHTML = models.map(row => `
  <tr>
    <td class="px-3 py-1 gap-1 font-semibold"><span class="${row.seriesClass} px-2 p-1 rounded-lg inline-block">${row.series}</span></td>
    <td class="px-3 py-1"><span class="${row.nameClass} px-2 p-1 rounded-lg inline-block">${row.name}</span></td>
    <td class="px-3 py-1"><span class="${row.typeClass} px-2 p-1 rounded-lg inline-block" data-i18n="${row.type}"></span></td>
    <td class="px-3 py-1"><span class="${row.domainClass} px-2 p-1 rounded-lg inline-block" data-i18n="${row.domain}"></span></td>
    <td class="px-3 py-1"><span class="${row.speedClass} px-2 p-1 rounded-lg inline-block" data-i18n="${row.speed}"></span></td>
    <td class="px-3 py-1"><span class="${row.costClass} px-2 p-1 rounded-lg inline-block" data-i18n="${row.cost}"></span></td>
    <td class="px-3 py-1"><span class="${row.brandClass} px-2 p-1 rounded-lg inline-block">${row.brand}</span></td>
    <td class="px-3 py-1"><span class="${row.deployClass} px-2 p-1 rounded-lg inline-block" data-i18n="${row.deploy}"></span></td>
  </tr>
`).join('');