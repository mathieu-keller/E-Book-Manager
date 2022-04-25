<script setup lang="ts">
import Button from '@/UI/Button.vue';
import Upload from '@/Upload/Upload.vue';
import router from "@/router";
import {ref} from "vue-demi";
import {onMounted, onUnmounted, watch} from "vue";
import {ApplicationStore} from '@/stores/ApplicationStore';
import type {LocationQueryValue} from "vue-router";

let isDarkMode = ref<boolean>(window.matchMedia('(prefers-color-scheme: dark)').matches);

const setDarkClass = () => {
  if (isDarkMode.value) {
    document.documentElement.classList.add('dark');
  } else {
    document.documentElement.classList.remove('dark');
  }
};

const setDark = (): void => {
  isDarkMode.value = !isDarkMode.value;
  setDarkClass();
};

onMounted(() => {
  setDarkClass();
});
const uploadFile = ref<boolean>(false);
const setUploadFile = (value: boolean) => {
  uploadFile.value = value;
};

const querySearch = ref<LocationQueryValue>(null);
const watchCleaner = watch(router.currentRoute, (newRoute, _) => {
  if (!Array.isArray(newRoute.query.q)) {
    querySearch.value = newRoute.query.q;
  }
});


onUnmounted(() => {
  watchCleaner();
});
const store = ApplicationStore();
const onInput = (inputEvent: Event) => {
  const target = inputEvent.target as HTMLInputElement;
  if (target.value === "") {
    router.push("/");
  } else {
    router.push(`/search?q=${target.value}`);
  }
};

onMounted(() => {
  store.$subscribe((_, store) => document.title = `E-Book: ${store.headerText}`);
});

</script>

<template>
  <Upload v-if="uploadFile" v-bind="{onClose: () => setUploadFile(false)}"/>
  <div class="flex flex-row justify-between border-b-2">
    <div>
      <Button button-type="default" button-text="Home" v-bind="{onClick: () => router.push('/')}"/>
      <Button button-type="default" v-bind="{onClick: setDark, buttonText: isDarkMode ? 'Light mode' : 'Dark mode'}"/>
    </div>
    <h1 class="text-5xl m-2 font-bold break-all">{{ store.headerText }}</h1>
    <Button button-type="primary" button-text="Upload!" v-bind="{onClick: ()=> setUploadFile(true)}"/>
  </div>
  <input
      class="w-[100%] text-5xl bg-slate-300 dark:bg-slate-700"
      placeholder="Search Books, Authors and Subjects"
      v-on:input="onInput"
      v-bind="{
        value: querySearch,
      }"
  />
</template>

