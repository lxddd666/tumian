import { http, jumpExport } from '@/utils/http/axios';

// 获取财务指标分析表列表
export function List(params) {
  return http.request({
    url: '/financialIndicators/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除财务指标分析表
export function Delete(params) {
  return http.request({
    url: '/financialIndicators/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑财务指标分析表
export function Edit(params) {
  return http.request({
    url: '/financialIndicators/edit',
    method: 'POST',
    params,
  });
}

// 获取财务指标分析表指定详情
export function View(params) {
  return http.request({
    url: '/financialIndicators/view',
    method: 'GET',
    params,
  });
}

// 导出财务指标分析表
export function Export(params) {
  jumpExport('/financialIndicators/export', params);
}