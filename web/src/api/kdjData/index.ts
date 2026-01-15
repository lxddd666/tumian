import { http, jumpExport } from '@/utils/http/axios';

// 获取KDJ随机指标数据表列表
export function List(params) {
  return http.request({
    url: '/kdjData/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除KDJ随机指标数据表
export function Delete(params) {
  return http.request({
    url: '/kdjData/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑KDJ随机指标数据表
export function Edit(params) {
  return http.request({
    url: '/kdjData/edit',
    method: 'POST',
    params,
  });
}

// 获取KDJ随机指标数据表指定详情
export function View(params) {
  return http.request({
    url: '/kdjData/view',
    method: 'GET',
    params,
  });
}

// 导出KDJ随机指标数据表
export function Export(params) {
  jumpExport('/kdjData/export', params);
}