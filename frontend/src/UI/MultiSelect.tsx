import { createSignal, For, Show } from 'solid-js';

// eslint-disable-next-line no-unused-vars, no-use-before-define
type MultiSelectProps<T extends { [A in R]: string }, R extends keyof T> = {
  readonly data: readonly T[];
  readonly showValue: R;
  readonly onChange: (data: T[]) => void;
  readonly selected: readonly T[];
}

// eslint-disable-next-line no-unused-vars, no-use-before-define
function MultiSelect<T extends { [A in R]: string }, R extends keyof T> (props: MultiSelectProps<T, R>) {
  const [value, setValue] = createSignal<string | null>(null);
  const [showDataSet, setShowDataSet] = createSignal<boolean>(false);

  const selectData = (data: T) => {
    props.onChange([...props.selected, data]);
  };

  return (
    <div class="relative w-[100%]" onMouseLeave={() => setShowDataSet(false)}>
      <input
        value={value() || ''}
        onFocusIn={() => setShowDataSet(true)}
        onInput={e => setValue(e.currentTarget.value)}
        class="w-[100%] text-xl bg-slate-300 dark:bg-slate-700"
        onKeyUp={(e) => {
          e.preventDefault();
          const filteredData = props.data
            .filter(d => props.selected.find((select) => select[props.showValue] === d[props.showValue]) === undefined)
            .filter(d => d[props.showValue].toLowerCase().startsWith(value()?.toLowerCase() || ''));
          if (e.keyCode === 13 && filteredData.length > 0) {
            selectData(filteredData[0]);
          }
        }}
      />
      <Show when={showDataSet()}>
        <div class="absolute w-[100%] border-2 border-white dark:bg-slate-900 dark:text-slate-300 bg-slate-50 text-slate-800 z-10">
          <For each={props.data
            .filter(d => props.selected.find((select) => select[props.showValue] === d[props.showValue]) === undefined)
            .filter(d => d[props.showValue].toLowerCase().startsWith(value()?.toLowerCase() || ''))}>
            {(d) => <p
              class="hover:text-white hover:bg-slate-400 hover:border-transparent dark:hover:bg-slate-500 cursor-pointer"
              onClick={() => selectData(d)}
            >
              {d[props.showValue]}
            </p>}
          </For>
        </div>
      </Show>
    </div>
  );
}

export default MultiSelect;
