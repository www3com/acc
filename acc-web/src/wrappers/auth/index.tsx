import {Redirect} from 'umi'

export default (props: any) => {
  // const {isLogin} = useAuth();
  if (true) {
    return <div>{props.children}</div>;
  } else {
    return <Redirect to="/sign-in"/>;
  }
}
