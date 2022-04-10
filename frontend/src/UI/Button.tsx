import React from 'react';

type ButtonProps = {
  readonly children?: React.ReactNode;
} & React.DetailedHTMLProps<React.ButtonHTMLAttributes<HTMLButtonElement>, HTMLButtonElement>
export const Button = (props: ButtonProps) => {
  return (
    <button {...props}
            className="bg-transparent dark:border-slate-200 dark:hover:bg-slate-500 dark:text-slate-50 border-slate-500 hover:bg-slate-400 text-slate-800 font-semibold hover:text-white py-2 px-4 border hover:border-transparent rounded"
    >
      {props.children}
    </button>
  );
};

export const PrimaryButton = (props: ButtonProps) => {
  return (
    <button {...props}
            className="dark:bg-red-900 dark:hover:bg-red-800 bg-red-500 hover:bg-red-400 text-white font-bold py-2 px-4 rounded">
      {props.children}
    </button>
  );
};

export default Button;
