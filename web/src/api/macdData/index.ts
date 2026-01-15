import { http, jumpExport } from '@/utils/http/axios';

// 获取MACD指标数据表列表
export function List(params) {
  return http.request({
    url: '/macdData/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除MACD指标数据表
export function Delete(params) {
  return http.request({
    url: '/macdData/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑MACD指标数据表
export function Edit(params) {
  return http.request({
    url: '/macdData/edit',
    method: 'POST',
    params,
  });
}

// 获取MACD指标数据表指定详情
export function View(params) {
  return http.request({
    url: '/macdData/view',
    method: 'GET',
    params,
  });
}

// 导出MACD指标数据表
export function Export(params) {
  jumpExport('/macdData/export', params);
}