import {defineStore} from 'pinia';

export const ApplicationStore = defineStore<'application', { headerText: string }, {}, {
  setHeaderText: (text: string) => void;
}>({
  id: 'application',
  state: () => ({
    headerText: 'Manager'
  }),
  actions: {
    setHeaderText(text: string) {
      this.headerText = text;
    }
  }
});
