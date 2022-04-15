import React from 'react';
type BadgeProps = {
  readonly text: string;
  readonly onClick?: ()=>void;
}
const Badge = (props: BadgeProps) => {
  let classNames = "border-2 p-2 w-max m-2 rounded-2xl bg-transparent dark:border-slate-200 dark:text-slate-50 border-slate-500 text-slate-800 font-semibold";
  if(props.onClick !== undefined) {
    classNames += " hover:text-white hover:bg-slate-400 hover:border-transparent dark:hover:bg-slate-500 cursor-pointer"
  }
  return (
    <div onClick={props.onClick} className={classNames}>
      {props.text}
    </div>
  );
};

export default Badge;
