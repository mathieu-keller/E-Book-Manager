import {Component} from "solid-js";

type BadgeProps = {
  readonly onClick?: () => void;
  readonly text: string;
}

const Badge: Component<BadgeProps> = (props) => {
  const hoverClasses = props.onClick ? "hover:text-white hover:bg-slate-400 hover:border-transparent dark:hover:bg-slate-500 cursor-pointer" : "";
  return (
    <div
      onClick={props.onClick}
      class={"float-left border-2 p-2 w-max m-2 rounded-2xl bg-transparent dark:border-slate-200 dark:text-slate-50 border-slate-500 text-slate-800 font-semibold " + hoverClasses}
    >
      {props.text}
    </div>
  );
};

export default Badge;
