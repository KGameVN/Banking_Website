import LoginWrapper from "./wrappers/loginwarrpers"; // chỉ import wrapper
import style from './styles/page.module.css';

export default function Page() {
  return (
    <div>
      <nav className={style.container}>
        <span>Welcome To World Bank</span>
          <LoginWrapper /></nav>
      <div className={style.mainview}></div>
    </div>
  );
}
