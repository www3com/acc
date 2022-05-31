import {Redirect} from 'umi'
import {getToken} from "@/components/token";

export default ({children}: any) => {
  if (getToken() == null) {
    return <Redirect to="/sign-in"/>;
  } else {
    return <div>{children}</div>;
  }
}
