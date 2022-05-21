import Modal from '../UI/Modal';
import { Component, createSignal, For, onMount } from 'solid-js';
import { BookType } from '../Book/Book.type';
import Rest from '../Rest';
import { BOOK_API } from '../Api/Api';
import Badge from '../UI/Badge';
import { Subject } from './Subject.type';
import MultiSelect from '../UI/MultiSelect';
import { Button, PrimaryButton } from '../UI/Button';

type EditSubjectsProps = {
  readonly onClose: () => void;
  readonly title: string;
};

const EditSubjects: Component<EditSubjectsProps> = (props) => {
  const [subjects, setSubjects] = createSignal<Subject[]>([]);
  const [allSubjects, setAllSubjects] = createSignal<Subject[]>([]);
  onMount(() => {
    Rest.get<BookType>(BOOK_API(props.title))
      .then(r => setSubjects(r.data.subjects));
    Rest.get<Subject[]>('/api/subjects').then(r => setAllSubjects(r.data));
  });

  return (
    <Modal
      onClose={props.onClose}
      title={props.title}
    >
      <div class="p-5">
        <h1 class="font-bold text-2xl">Subjects:</h1>
        <MultiSelect<Subject, 'name'> data={allSubjects()} selected={subjects()} showValue="name" onChange={data => setSubjects(data)}/>
        <div class="flex flex-row flex-wrap">
          <For each={subjects()}>
            {(subject) => <Badge text={subject.name} onRemove={() => {
              setSubjects(subjects()!.filter(sub => sub.name !== subject.name));
            }}/>}
          </For>
        </div>
      </div>
      <footer class="border-t-2 w-full pt-5 flex justify-center">
        <div class="flex justify-around w-full">
          <PrimaryButton type="submit" form="upload-epub" className="w-1/2">
            Submit
          </PrimaryButton>
          <Button button-type="default" onClick={props.onClose} className="w-1/2">
            Close
          </Button>
        </div>
      </footer>
    </Modal>
  );
};

export default EditSubjects;
