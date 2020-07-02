import { FormGroup, ValidationErrors, ValidatorFn } from '@angular/forms';

export function getPasswordValidator(passwordControlName: string, verifyControlName: string): ValidatorFn {
  return (form: FormGroup): ValidationErrors | null => {
    const passwordControl = form.get(passwordControlName);
    const verifyControl = form.get(verifyControlName);
    return passwordControl && verifyControl && passwordControl.value !== verifyControl.value ? { passwordVerification: true } : null;
  };
}
