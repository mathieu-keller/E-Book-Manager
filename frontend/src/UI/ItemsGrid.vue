<script setup lang="ts">
import ItemCard from './ItemCard.vue';
import type {LibraryItemType} from "@/Library/LibraryItem.type";

defineProps<{
  items: LibraryItemType[];
  onClick: (item: LibraryItemType) => void;
}>();
</script>
<template>
  <div class="flex flex-wrap flex-row justify-center">
    <ItemCard
        v-for="item in items"
        v-bind="{
            name: item.title,
            cover: item.cover,
            itemCount: item.itemType === 'collection' ? item.bookCount : undefined,
            id: item.id,
            itemType: item.itemType,
            onClick: () => onClick(item)
        }"
    />
  </div>
</template>
