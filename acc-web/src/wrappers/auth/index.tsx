import {Redirect} from 'umi'
import {getToken} from "@/components/token";

export default (props: any) => {
  if (getToken() == null) {
    return <Redirect to="/sign-in"/>;
  } else {
    return <div>{props.children}</div>;
  }
}
