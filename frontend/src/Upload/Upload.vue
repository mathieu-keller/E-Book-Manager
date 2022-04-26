<script setup lang="ts">
import Modal from '../UI/Modal.vue';
import Button from "@/UI/Button.vue";
import Rest from "@/Rest";
import {ref} from "vue-demi";
import {UPLOAD_API} from "@/api/Api";
import upload_icon from '@/assets/upload.svg';

const props = defineProps<{
  onClose: () => void
}>();

const maxSize = ref<number | null>(null);
const current = ref<number | null>(null);

const uploadBooks = async (data: FormData): Promise<void> => {
  const response = await Rest.post(UPLOAD_API, data, {
    onUploadProgress: (e: ProgressEvent): void => {
      maxSize.value = e.total;
      current.value = e.loaded;
    }
  });
  if (response.status === 200) {
    location.reload();
  }
};

const onSubmit = (e: any): void => {
  e.preventDefault();
  const form = new FormData(e.currentTarget);
  uploadBooks(form)
      .then((): void => props.onClose())
      .catch((e: string): void => console.error(e));
};

</script>
<template>
  <Modal
      v-bind="{
      onClose: onClose
    }"
      title="Upload E-Book">
    <template #footer>
      <div class="flex justify-around w-full">
        <Button button-type="primary" type="submit" form="upload-epub">
          <img
              class="dark:invert invert-0 h-8 mr-1"
              v-bind="{
              src: upload_icon
            }"
              alt="upload"
          /> Upload
        </Button>
        <Button button-type="default" v-bind="{
          onClick: onClose
        }">
          Close
        </Button>
      </div>
    </template>
    <template #default>
      <form
          id="upload-epub"
          @submit="onSubmit"
      >
        <input type="file" accept="application/epub+zip" name="myFiles" multiple/>
      </form>
      <div v-if="current !== null && maxSize !== null">
        <progress v-bind="{value: current, max: maxSize}"/>
        {{ (Math.round((current / maxSize) * 10000)) / 100 }}% <br/>
        ({{ current }} / {{ maxSize }})
      </div>
    </template>
  </Modal>

</template>
