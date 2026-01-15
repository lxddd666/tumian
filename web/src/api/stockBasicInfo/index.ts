import { http, jumpExport } from '@/utils/http/axios';

// 获取股票基础信息表列表
export function List(params) {
  return http.request({
    url: '/stockBasicInfo/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除股票基础信息表
export function Delete(params) {
  return http.request({
    url: '/stockBasicInfo/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑股票基础信息表
export function Edit(params) {
  return http.request({
    url: '/stockBasicInfo/edit',
    method: 'POST',
    params,
  });
}

// 获取股票基础信息表指定详情
export function View(params) {
  return http.request({
    url: '/stockBasicInfo/view',
    method: 'GET',
    params,
  });
}

// 导出股票基础信息表
export function Export(params) {
  jumpExport('/stockBasicInfo/export', params);
}