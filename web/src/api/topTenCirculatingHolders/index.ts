import { http, jumpExport } from '@/utils/http/axios';

// 获取公司十大流通股东表 (数据来源于定期报告)[citation:4]列表
export function List(params) {
  return http.request({
    url: '/topTenCirculatingHolders/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除公司十大流通股东表 (数据来源于定期报告)[citation:4]
export function Delete(params) {
  return http.request({
    url: '/topTenCirculatingHolders/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑公司十大流通股东表 (数据来源于定期报告)[citation:4]
export function Edit(params) {
  return http.request({
    url: '/topTenCirculatingHolders/edit',
    method: 'POST',
    params,
  });
}

// 获取公司十大流通股东表 (数据来源于定期报告)[citation:4]指定详情
export function View(params) {
  return http.request({
    url: '/topTenCirculatingHolders/view',
    method: 'GET',
    params,
  });
}

// 导出公司十大流通股东表 (数据来源于定期报告)[citation:4]
export function Export(params) {
  jumpExport('/topTenCirculatingHolders/export', params);
}