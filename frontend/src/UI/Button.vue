<script setup lang="ts">
defineProps<{
  buttonType: 'primary' | 'link' | 'default';
  buttonText: string;
  href?: string;
  download?: string;
  className?: string;
  onClick?: () => void;
  type?: "button" | "reset" | "submit";
  form?: string;
}>();
</script>
<template>
  <button
      v-if="buttonType === 'default'"
      v-on:click="onClick"
      v-bind="{
        type,
        form
      }"
      @click="{onClick}"
      class="bg-transparent dark:border-slate-200 dark:hover:bg-slate-500 dark:text-slate-50 border-slate-500 hover:bg-slate-400 text-slate-800 font-semibold hover:text-white py-2 px-4 border hover:border-transparent rounded"
  >
    {{ buttonText }}
  </button>
  <button
      v-if="buttonType === 'primary'"
      v-on:click="onClick"
      v-bind="{
        type,
        form
      }"
      class="dark:bg-red-900 dark:hover:bg-red-800 bg-red-500 hover:bg-red-400 text-white font-bold py-2 px-4 rounded"
  >
    {{ buttonText }}
  </button>
  <a
      v-bind="{
        href,
        download: download,
        class: 'w-[100%] text-center cursor-pointer bg-transparent dark:border-slate-200 dark:hover:bg-slate-500 dark:text-slate-50 border-slate-500 hover:bg-slate-400 text-slate-800 font-semibold hover:text-white py-2 px-4 border hover:border-transparent rounded'
        + ' '+(className? className : '')
    }"
      v-if="buttonType === 'link'"
  >
    {{ buttonText }}
  </a>
</template>
