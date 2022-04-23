<script setup lang="ts">
import Badge from '@/UI/Badge.vue';
import Button from '@/UI/Button.vue';
import type {BookType} from "@/Book/Book.type";
import router from "@/router";
import defaultCover from '../assets/cover.jpg';
import {ref} from "vue-demi";
import Rest from "../Rest";
import {onMounted} from "vue";
import {ApplicationStore} from "@/stores/ApplicationStore";

const title: string = router.currentRoute.value.params.title as string;
const book = ref<BookType>();
const getBook = async (): Promise<BookType> => {
  const response = await Rest.get<BookType>(`/api/book/${title}`);
  return response.data;
};

const store = ApplicationStore();
onMounted(()=> {
  store.setHeaderText(title);
  getBook().then(r => book.value = r);
});

</script>


<template>
  <div class="mt-10 flex justify-center">
    <div class="grid max-w-[80%]" v-if="book !== undefined">
      <img
          v-bind="{
        src: book.cover !== null ? `data:image/jpeg;base64,${book.cover}` : defaultCover,
        alt: `cover picture of ${book.title}`
        }"
      />
      <div class="grid-cols-1 grid h-max">
        <div class="m-5">
          <h1>Authors:</h1>
          <Badge
              v-for="author in book.authors"
              v-bind="{
               onClick: () => router.push(`/search?q=${author.name}`),
               text: author.name
             }"
          />
        </div>
        <div class="m-5">
          <h1>Subjects:</h1>
          <Badge
              v-for="subject in book.subjects"
              v-bind="{
               onClick: () => router.push(`/search?q=${subject.name}`),
               text: subject.name
             }"
          />
        </div>
      </div>
      <div class="col-start-1 col-end-3 mt-5 flex justify-self-stretch">
        <Button
            button-text="Download"
            v-bind="{
              href: `/download/${book.id}`,
              download: `${book.title}.epub`
            }"
            button-type="link"/>
      </div>
    </div>
  </div>
</template>

