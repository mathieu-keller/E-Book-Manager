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
const searchBooks = async (search: LocationQueryValue | LocationQueryValue[]): Promise<BookType[]> => {
  if(search === null || Array.isArray(search)){
    return Promise.reject();
  }
  const response = await Rest.get<BookType[]>(`/api/book?q=${search}`);
  return response.data;
};
const store = ApplicationStore();
let timer: null | number = null;
const watchCleaner = watch(router.currentRoute, (newRoute, _) => {
  if (newRoute.query.q !== undefined) {
    store.setHeaderText(`Search: ${newRoute.query.q}`);
    if (timer !== null) {
      window.clearTimeout(timer);
    }
    timer = window.setTimeout(() => searchBooks(newRoute.query.q).then(r => books.value = r), 500);
  }
});
const openItem = (book: BookType): void => {
  router.push(`/book/${book.title}`);
};

onMounted(() => {
  store.setHeaderText(`Search: ${router.currentRoute.value.query.q}`);
  searchBooks(router.currentRoute.value.query.q).then(r => books.value = r);
});

onUnmounted(() => {
  if (timer !== null) {
    window.clearTimeout(timer);
  }
  watchCleaner();
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
</template>
