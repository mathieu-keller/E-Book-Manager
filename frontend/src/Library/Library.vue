<script setup lang="ts">

import type {LibraryItemType} from './LibraryItem.type';
import ItemsGrid from '../UI/ItemsGrid.vue';
import Rest from "../Rest";
import router from "@/router";
import {onMounted, onUnmounted} from "vue";
import {ApplicationStore} from "@/stores/ApplicationStore";
import {type LibraryItemStoreType, LibraryStore} from "@/stores/LibraryStore";
import {ref, type UnwrapRef} from "vue-demi";


const libraryStore = LibraryStore();
const items = ref<LibraryItemType[]>(libraryStore.items);
const getLibraryItems = async (page: number): Promise<LibraryItemType[]> => {
  const response = await Rest.get<LibraryItemType[]>(`/api/all?page=${page}`);
  return response.data;
};

const openItem = (item: LibraryItemType) => {
  router.push(`/${item.itemType}/${item.title}`);
};

const loading = ref<boolean>(false);
const page = ref<number>(libraryStore.page);

const shouldLoadNextPage = (): void => {
  const element = document.querySelector('#loading-trigger');
  const position = element?.getBoundingClientRect();

  if (position !== undefined && !loading.value && position.top >= 0 && position.bottom <= window.innerHeight) {
    loading.value = true;
    getLibraryItems(page.value).then(r => {
      if (r.length > 0) {
        libraryStore.addAll(r);
        libraryStore.setPage(page.value + 1);
        loading.value = false;
        window.setTimeout(() => shouldLoadNextPage(), 0);
      } else if (r.length === 0 || r.length > 32) {
        libraryStore.setAllLoaded(true);
      }
    });
  }
};

const allLoaded = ref<boolean>(libraryStore.allItemsLoaded);
const applicationStore = ApplicationStore();
onMounted(() => {
  applicationStore.$reset();
  libraryStore.$subscribe((_,
                           state: UnwrapRef<LibraryItemStoreType>) => {
    items.value = state.items;
    allLoaded.value = state.allItemsLoaded;
    page.value = state.page;
  });
  window.addEventListener('scroll', shouldLoadNextPage);
  shouldLoadNextPage();
});

onUnmounted(() => {
  window.removeEventListener('scroll', shouldLoadNextPage);
});

</script>

<template>
  <ItemsGrid v-bind="{
        onClick: (item) => openItem(item),
        items: items
    }"/>
  <div v-if="!allLoaded && !libraryStore.allItemsLoaded" id="loading-trigger" class="m-5 text-center text-5xl">Loading....</div>
</template>
