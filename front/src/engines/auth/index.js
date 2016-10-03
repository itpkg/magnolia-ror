import Confirm from './Confirm'
import ForgotPassword from './ForgotPassword'
import SignIn from './SignIn'

export default {
  routes: [
    {path: '/users/sign-in', component: SignIn},
    {path: '/users/confirm', component: Confirm},
    {path: '/users/forgot-password', component: ForgotPassword}
  ]
}
