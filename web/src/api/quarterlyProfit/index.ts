import { http, jumpExport } from '@/utils/http/axios';

// 获取季度利润数据表 (近一年各季度)列表
export function List(params) {
  return http.request({
    url: '/quarterlyProfit/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除季度利润数据表 (近一年各季度)
export function Delete(params) {
  return http.request({
    url: '/quarterlyProfit/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑季度利润数据表 (近一年各季度)
export function Edit(params) {
  return http.request({
    url: '/quarterlyProfit/edit',
    method: 'POST',
    params,
  });
}

// 获取季度利润数据表 (近一年各季度)指定详情
export function View(params) {
  return http.request({
    url: '/quarterlyProfit/view',
    method: 'GET',
    params,
  });
}

// 导出季度利润数据表 (近一年各季度)
export function Export(params) {
  jumpExport('/quarterlyProfit/export', params);
}