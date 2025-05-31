//

"user client";

import React from "react";

type ButtonProps = {
  onClick: () => void;
  label: string;
  className: string;
}

export default function Button({onClick, label, className}: ButtonProps){
  return (
    <button
      onClick={onClick}
      className={className}
    >{label}
    </button>
  );
}
