import { http, jumpExport } from '@/utils/http/axios';

// 获取公司股东户数统计表 (按报告期统计)列表
export function List(params) {
  return http.request({
    url: '/shareholderCount/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除公司股东户数统计表 (按报告期统计)
export function Delete(params) {
  return http.request({
    url: '/shareholderCount/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑公司股东户数统计表 (按报告期统计)
export function Edit(params) {
  return http.request({
    url: '/shareholderCount/edit',
    method: 'POST',
    params,
  });
}

// 获取公司股东户数统计表 (按报告期统计)指定详情
export function View(params) {
  return http.request({
    url: '/shareholderCount/view',
    method: 'GET',
    params,
  });
}

// 导出公司股东户数统计表 (按报告期统计)
export function Export(params) {
  jumpExport('/shareholderCount/export', params);
}