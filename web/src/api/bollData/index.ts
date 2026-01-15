import { http, jumpExport } from '@/utils/http/axios';

// 获取布林带(BOLL)指标数据表列表
export function List(params) {
  return http.request({
    url: '/bollData/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除布林带(BOLL)指标数据表
export function Delete(params) {
  return http.request({
    url: '/bollData/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑布林带(BOLL)指标数据表
export function Edit(params) {
  return http.request({
    url: '/bollData/edit',
    method: 'POST',
    params,
  });
}

// 获取布林带(BOLL)指标数据表指定详情
export function View(params) {
  return http.request({
    url: '/bollData/view',
    method: 'GET',
    params,
  });
}

// 导出布林带(BOLL)指标数据表
export function Export(params) {
  jumpExport('/bollData/export', params);
}