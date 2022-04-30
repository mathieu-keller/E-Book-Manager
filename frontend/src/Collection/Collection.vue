<script setup lang="ts">
import router from "@/router";
import type {CollectionType} from "@/Collection/Collection.type";
import type {BookType} from "@/Book/Book.type";
import type {LibraryItemType} from "@/Library/LibraryItem.type";
import Rest from "../Rest";
import {ref, type UnwrapRef} from "vue-demi";
import type {Ref} from "vue-demi";
import ItemsGrid from "../UI/ItemsGrid.vue";
import {onMounted} from "vue";
import {ApplicationStore} from "@/stores/ApplicationStore";
import {CollectionStore, type CollectionStoreType} from "@/stores/CollectionStore";
import type {SubscriptionCallbackMutation} from "pinia";
import {COLLECTION_API} from "@/api/Api";

const title = router.currentRoute.value.params.title as string;

const getCollection = async (): Promise<CollectionType> => {
  const response = await Rest.get<CollectionType>(COLLECTION_API(title));
  return response.data;
};

const collectionStore = CollectionStore();

const collection: Ref<BookType[]> | Ref<null> = ref(collectionStore.collections[title] || null);


const openItem = (item: LibraryItemType) => {
  router.push(`/${item.itemType}/${encodeURIComponent(item.title)}`);
};
const applicationStore = ApplicationStore();
onMounted(() => {
  applicationStore.setHeaderText(title);
  collectionStore.$subscribe((mutation: SubscriptionCallbackMutation<CollectionStoreType>,
                              state: UnwrapRef<CollectionStoreType>) => collection.value = state.collections[title]);
  if (collection.value === null) {
    getCollection().then(r => collectionStore.set(r.title, r.books));
  }
});

</script>

<template>
  <ItemsGrid v-if="collection !== null" v-bind="{
        onClick: (item) => openItem(item),
        items: collection === null? [] : collection.map((book: BookType) => ({
          id: book.id,
          cover: book.cover,
          title: book.title,
          itemType: 'book',
          bookCount: 1
        }))
    }"/>
</template>
