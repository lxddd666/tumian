import { http, jumpExport } from '@/utils/http/axios';

// 获取基金持股明细表 (来源于基金定期报告)列表
export function List(params) {
  return http.request({
    url: '/fundStockHolding/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除基金持股明细表 (来源于基金定期报告)
export function Delete(params) {
  return http.request({
    url: '/fundStockHolding/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑基金持股明细表 (来源于基金定期报告)
export function Edit(params) {
  return http.request({
    url: '/fundStockHolding/edit',
    method: 'POST',
    params,
  });
}

// 获取基金持股明细表 (来源于基金定期报告)指定详情
export function View(params) {
  return http.request({
    url: '/fundStockHolding/view',
    method: 'GET',
    params,
  });
}

// 导出基金持股明细表 (来源于基金定期报告)
export function Export(params) {
  jumpExport('/fundStockHolding/export', params);
}