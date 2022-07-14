// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { FormDescriptor } from './FormDescriptor';

export type FormProps = {
    form: FormDescriptor;
    children: React.ReactNode;
};

export const Form = (props: FormProps): JSX.Element => {
    const { submitAction, submitMethod, csrfToken } = props.form

    return (
        <form
            method={submitMethod}
            action={submitAction}>
            {
                csrfToken !== undefined && csrfToken !== '' &&
                <input type="hidden" name="csrf_token" value={csrfToken} />
            }
            {props.children}
        </form>
    );
};
