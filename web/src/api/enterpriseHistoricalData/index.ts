import { http, jumpExport } from '@/utils/http/axios';

// 获取企业级历史行情数据表 (K线数据)列表
export function List(params) {
  return http.request({
    url: '/enterpriseHistoricalData/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除企业级历史行情数据表 (K线数据)
export function Delete(params) {
  return http.request({
    url: '/enterpriseHistoricalData/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑企业级历史行情数据表 (K线数据)
export function Edit(params) {
  return http.request({
    url: '/enterpriseHistoricalData/edit',
    method: 'POST',
    params,
  });
}

// 获取企业级历史行情数据表 (K线数据)指定详情
export function View(params) {
  return http.request({
    url: '/enterpriseHistoricalData/view',
    method: 'GET',
    params,
  });
}

// 导出企业级历史行情数据表 (K线数据)
export function Export(params) {
  jumpExport('/enterpriseHistoricalData/export', params);
}