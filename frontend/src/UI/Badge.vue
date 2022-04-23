<script setup lang="ts">
defineProps<{
  text: string;
  onClick?: () => void;
}>();
</script>

<template>
  <div v-if="!onClick"
       class="float-left border-2 p-2 w-max m-2 rounded-2xl bg-transparent dark:border-slate-200 dark:text-slate-50 border-slate-500 text-slate-800 font-semibold">
    {{text}}
  </div>

  <div v-if="onClick" @click="onClick"
       class="float-left border-2 p-2 w-max m-2 rounded-2xl bg-transparent dark:border-slate-200 dark:text-slate-50 border-slate-500 text-slate-800 font-semibold  hover:text-white hover:bg-slate-400 hover:border-transparent dark:hover:bg-slate-500 cursor-pointer">
    {{text}}
  </div>
</template>
