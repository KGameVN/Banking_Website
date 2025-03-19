import Image from "next/image";
import LoginWrapper from "./wrappers/loginwarrpers"; // chỉ import wrapper
import a from './styles/page.module.css';

export default function Page() {
  return (
    <div>
      <nav className={a.container}>Navbar</nav>
      <div className="flex flex-col items-center justify-center p-10">
        <Image src="/logo.png" alt="Logo" width={200} height={200} />
        <LoginWrapper />
      </div>
    </div>
  );
}
