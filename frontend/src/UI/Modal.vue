<script setup lang="ts">
defineProps<{
  readonly title: string;
  readonly onClose: () => void;
}>();
</script>


<template>
  <div @click="onClose" class="fixed top-0 left-0 right-0 bottom-0 z-40 bg-gray-700/75"/>
  <div
      class="dark:bg-slate-900 dark:text-slate-300 bg-slate-50 text-slate-800 z-50 flex flex-col fixed max-h-[50%] w-2/4 left-1/4 top-1/4 my-auto flex justify-center items-center">
    <h1 class="border-b-2 w-full flex justify-center">{{ title }}</h1>
    <div class="word-wrap-break-word max-h-[80%]  overflow-y-auto ">
      <slot name="default"></slot>
    </div>
    <footer class="border-t-2 w-full flex justify-center">
      <slot name="footer"></slot>
    </footer>
  </div>
</template>

