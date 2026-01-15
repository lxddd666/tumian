import { http, jumpExport } from '@/utils/http/axios';

// 获取股东户数变化记录表 (记录相邻报告期的户数变化)列表
export function List(params) {
  return http.request({
    url: '/shareholderChange/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除股东户数变化记录表 (记录相邻报告期的户数变化)
export function Delete(params) {
  return http.request({
    url: '/shareholderChange/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑股东户数变化记录表 (记录相邻报告期的户数变化)
export function Edit(params) {
  return http.request({
    url: '/shareholderChange/edit',
    method: 'POST',
    params,
  });
}

// 获取股东户数变化记录表 (记录相邻报告期的户数变化)指定详情
export function View(params) {
  return http.request({
    url: '/shareholderChange/view',
    method: 'GET',
    params,
  });
}

// 导出股东户数变化记录表 (记录相邻报告期的户数变化)
export function Export(params) {
  jumpExport('/shareholderChange/export', params);
}