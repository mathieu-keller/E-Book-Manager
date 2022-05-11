<script setup lang="ts">
defineProps<{
  buttonType: 'primary' | 'link' | 'default';
  href?: string;
  download?: boolean;
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
        form,
        class:'bg-transparent flex justify-center items-center dark:border-slate-200 dark:hover:bg-slate-500 dark:text-slate-50 border-slate-500 hover:bg-slate-400 text-slate-800 font-semibold hover:text-white py-1 px-2 border hover:border-transparent rounded'
        + ' '+(className? className : '')
  }"
      @click="{onClick}">
    <slot/>
  </button>
  <button
      v-if="buttonType === 'primary'"
      v-on:click="onClick"
      v-bind="{
        type,
        form
      }"
      class="dark:bg-red-900 flex justify-center items-center dark:hover:bg-red-800 bg-red-500 hover:bg-red-400 text-white font-bold py-1 px-2 rounded"
  >
    <slot/>
  </button>
  <a
      v-bind="{
        href,
        download: download,
        class: 'w-[100%] text-center flex justify-center items-center cursor-pointer bg-transparent dark:border-slate-200 dark:hover:bg-slate-500 dark:text-slate-50 border-slate-500 hover:bg-slate-400 text-slate-800 font-semibold hover:text-white py-1 px-2 border hover:border-transparent rounded'
        + ' '+(className? className : '')
    }"
      v-if="buttonType === 'link'"
  >
    <slot/>
  </a>
</template>
