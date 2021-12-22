import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Accounts from '../views/Accounts.vue'
import Payments from '../views/Payments.vue'
import Reports from '../views/Reports.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/payments',
    name: 'Payments',
    component: Payments
  },
  {
    path: '/accounts',
    name: 'Accounts',
    component: Accounts
  },
  {
    path: '/reports',
    name: 'reports',
    component: Reports
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
