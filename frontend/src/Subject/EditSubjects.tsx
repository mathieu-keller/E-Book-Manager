import Modal from '../UI/Modal';
import { Component, createSignal, For, onMount } from 'solid-js';
import { BookType } from '../Book/Book.type';
import Rest from '../Rest';
import { BOOK_API } from '../Api/Api';
import Badge from '../UI/Badge';
import { Subject } from './Subject.type';
import MultiSelect from '../UI/MultiSelect';

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
      <div class="flex flex-row">
        <MultiSelect<Subject, 'name'> data={allSubjects()} selected={subjects()} showValue="name" onChange={data => setSubjects(data)}/>
        <For each={subjects()}>
          {(subject) => <Badge text={subject.name} onRemove={() => {
            setSubjects(subjects()!.filter(sub => sub.name !== subject.name));
          }}/>}
        </For>
      </div>
    </Modal>
  );
};

export default EditSubjects;
