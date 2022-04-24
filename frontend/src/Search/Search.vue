<script setup lang="ts">
import type {BookType} from "@/Book/Book.type";
import ItemCard from '@/UI/ItemCard.vue';
import {ref, type Ref} from "vue-demi";
import Rest from "../Rest";
import router from "@/router";
import {onMounted, onUnmounted, watch} from "vue";
import {ApplicationStore} from "@/stores/ApplicationStore";
import type {LocationQueryValue} from "vue-router";

const books: Ref<BookType[]> = ref([]);
const loading = ref<boolean>(false);
const page = ref<number>(1);
const allLoaded = ref<boolean>(false);
const searchBooks = async (search: LocationQueryValue | LocationQueryValue[], currentPage: number): Promise<void> => {
  if (search === null || Array.isArray(search)) {
    return Promise.reject();
  }
  loading.value = true;
  const response = await Rest.get<BookType[]>(`/api/book?q=${search}&page=${currentPage}`);
  const data = response.data;
  if (data.length > 0) {
    if (currentPage === 1) {
      books.value = data;
      allLoaded.value = false;
      window.setTimeout(() => shouldLoadNextPage(), 50);
    } else {
      books.value = [...books.value, ...data];
    }
    page.value = currentPage + 1;
    loading.value = false;
  } else {
    allLoaded.value = true;
  }
};
const store = ApplicationStore();
let timer: null | number = null;
const watchCleaner = watch(router.currentRoute, (newRoute, _) => {
  if (newRoute.query.q !== undefined) {
    store.setHeaderText(`Search: ${newRoute.query.q}`);
    if (timer !== null) {
      window.clearTimeout(timer);
    }
    timer = window.setTimeout(() => searchBooks(newRoute.query.q, 1),
        500);
  }
});
const openItem = (book: BookType): void => {
  router.push(`/book/${book.title}`);
};

onMounted(() => {
  store.setHeaderText(`Search: ${router.currentRoute.value.query.q}`);
  searchBooks(router.currentRoute.value.query.q, 1);
});
const search = () => {
  searchBooks(router.currentRoute.value.query.q, page.value);
};
const shouldLoadNextPage = (): void => {
  const element = document.querySelector('#loading-trigger');
  const position = element?.getBoundingClientRect();

  if (position !== undefined && !loading.value && position.top >= 0 && position.bottom <= window.innerHeight) {
    search();
  }
};
window.addEventListener('scroll', shouldLoadNextPage);
onUnmounted(() => {
  if (timer !== null) {
    window.clearTimeout(timer);
  }
  watchCleaner();
  window.removeEventListener('scroll', shouldLoadNextPage);
});


</script>
<template>
  <div class="flex flex-wrap flex-row justify-center">
    <ItemCard
        v-for="book in books"
        v-bind="{
          name: book.title,
          cover: book.cover,
          id: book.id,
          onClick: () => openItem(book)
        }"
        item-type="book"
    />
  </div>
  <div @click="search" v-if="!allLoaded" id="loading-trigger"
       class="m-5 border cursor-pointer text-center text-5xl">Load More
  </div>
</template>
