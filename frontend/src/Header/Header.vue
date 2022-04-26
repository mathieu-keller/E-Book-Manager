<script setup lang="ts">
import Button from '@/UI/Button.vue';
import Upload from '@/Upload/Upload.vue';
import router from "@/router";
import {ref} from "vue-demi";
import {onMounted, onUnmounted} from "vue";
import {ApplicationStore} from '@/stores/ApplicationStore';
import upload_icon from '@/assets/upload.svg';

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

const store = ApplicationStore();
let search = ref<string>("");
let timer: null | number = null;

const clearTimer = () => {
  if (timer !== null) {
    window.clearTimeout(timer);
  }
};

const onInput = (inputEvent: Event) => {
  const target = inputEvent.target as HTMLInputElement;
  search.value = target.value;
  if (target.value === "") {
    clearTimer();
    router.push("/");
  } else {
    store.setHeaderText(`Search: ${search.value}`);
    if (timer !== null) {
      window.clearTimeout(timer);
    }
    timer = window.setTimeout(() => router.push(`/search?q=${encodeURIComponent(search.value)}`),
        500);
  }
};

onMounted(() => {
  store.$subscribe((_, store) => document.title = `E-Book: ${store.headerText}`);
  router.isReady().then(() => {
    if (!Array.isArray(router.currentRoute.value.query.q)) {
      search.value = router.currentRoute.value.query.q || "";
    }
  });
  router.afterEach(g => {
    if (search.value === "" && g.path === "/search" && !Array.isArray(g.query.q)) {
      search.value = g.query.q || "";
    } else if (search.value !== "" && g.path !== "/search") {
      search.value = "";
    }
  });
});

onUnmounted(() => {
  clearTimer();
});

</script>

<template>
  <Upload v-if="uploadFile" v-bind="{onClose: () => setUploadFile(false)}"/>
  <div class="flex flex-row justify-between border-b-2">
    <div>
      <Button button-type="default" v-bind="{onClick: () => router.push('/')}">
        Home
      </Button>
      <Button button-type="default" v-bind="{onClick: setDark}">
        {{ isDarkMode ? 'Light mode' : 'Dark mode' }}
      </Button>
    </div>
    <h1 class="text-5xl m-2 font-bold break-all">{{ store.headerText }}</h1>
    <Button button-type="primary" v-bind="{onClick: ()=> setUploadFile(true)}">
      <img
          class="dark:invert invert-0 h-8 mr-1"
          v-bind="{
              src: upload_icon
            }"
          alt="upload"
      /> Upload!
    </Button>
  </div>
  <input
      class="w-[100%] text-5xl bg-slate-300 dark:bg-slate-700"
      placeholder="Search Books, Authors and Subjects"
      v-on:input="onInput"
      v-bind="{
        value: search,
      }"
  />
</template>

