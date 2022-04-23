import {createRouter, createWebHistory} from 'vue-router';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'library',
      component: () => import('@/Library/Library.vue')
    }, {
      path: '/collection/:title',
      name: 'collection',
      component: () => import('@/Collection/Collection.vue')
    }
    , {
      path: '/book/:title',
      name: 'book',
      component: () => import('@/Book/Book.vue')
    }
    , {
      path: '/search',
      name: 'search',
      component: () => import('@/Search/Search.vue')
    }
  ]
});

export default router;
