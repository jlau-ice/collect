import type {RouteRecordRaw} from "vue-router";

export const menus: Array<RouteRecordRaw> = [
  {
    path: 'dashboard',
    name: 'Dashboard',
    component: () => import('@/views/Dashboard.vue'),
    meta: {title: '统计看板'}
  },
  {
    path: 'departments',
    name: 'Departments',
    component: () => import('@/views/Departments.vue'),
    meta: {title: '部门管理'}
  },
  {
    path: 'users',
    name: 'Users',
    component: () => import('@/views/Users.vue'),
    meta: {title: '人员管理'}
  },
  {
    path: 'tasks',
    name: 'Tasks',
    component: () => import('@/views/Tasks.vue'),
    meta: {title: '转发任务'}
  },
  {
    path: 'upload',
    name: 'Upload',
    component: () => import('@/views/Upload.vue'),
    meta: {title: '文件上传'}
  }
]