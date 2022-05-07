<script setup lang="ts">
import Button from "./Button.vue";
import defaultCover from '../assets/cover.jpg';
import {DOWNLOAD_API, DOWNLOAD_ORIGINAL_API} from "@/api/Api";
import {ref, type Ref} from "vue-demi";

defineProps<{
  itemCount?: number;
  cover: string | null;
  name: string;
  onClick: () => void;
  itemType: 'book' | 'collection';
  id: number;
}>();

const showOptions: Ref<boolean> = ref(false);
const setShowOptions = (value: boolean) => {
  showOptions.value = value;
};


</script>
<template>
  <div class="m-3 p-2 flex h-max w-80 flex-col">
    <div @click="onClick" class="flex justify-center hover:pb-3 cursor-pointer hover:mt-0 hover:mb-3 p-0 my-3 relative">

      <div v-if="itemCount !== undefined && itemCount !== null" class="absolute p-3 left-5 top-0 text-5xl bg-red-700 text-white rounded-b-full">
        {{ itemCount }}
      </div>

      <img
          v-bind="{
            src: cover === null? defaultCover : `data:image/jpeg;base64,${cover}`,
            alt: `cover picture of ${name}`
          }"
          width="270"
          height="470"
      />
    </div>
    <div>
      <h1
          @click="onClick"
          v-bind="{class: 'cursor-pointer text-center break-words text-2xl font-bold ' +
            (itemType === 'book' ? 'float-left w-10/12' : 'w-12/12')}"
      >
        {{ name }}
      </h1>
      <Button
          v-if="itemType === 'book'"
          button-type="default"
          class-name="w-2/12 float-right"
          v-bind="{
            onClick: () => setShowOptions(true)
          }"
      >
        ...
        <div v-if="showOptions" @mouseleave="() => setShowOptions(false)"
             class="absolute border-2 border-white dark:bg-slate-900 dark:text-slate-300 bg-slate-50 text-slate-800 z-10">
          <Button
              v-bind="{
                 download: true
              }"
              v-bind:href="DOWNLOAD_API(id)"
              button-type="link"
          >
            Download Book
          </Button>
          <Button
              v-bind="{
                 download: true
              }"
              v-bind:href="DOWNLOAD_ORIGINAL_API(id)"
              button-type="link"
          >
            Download Original Book
          </Button>
        </div>
      </Button>
    </div>
  </div>
</template>
