import {createRouter, createWebHistory} from 'vue-router';
import Library from '@/Library/Library.vue';
import Collection from "@/Collection/Collection.vue";
import Book from "@/Book/Book.vue";
import Search from "@/Search/Search.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'library',
      component: Library
    }, {
      path: '/collection/:title',
      name: 'collection',
      component: Collection
    }
    , {
      path: '/book/:title',
      name: 'book',
      component: Book
    }
    , {
      path: '/search',
      name: 'search',
      component: Search
    }
  ]
});

export default router;
