// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { FormDescriptor } from './FormDescriptor';

export type FormProps = {
    form: FormDescriptor;
    children: React.ReactNode;
};

export const Form = (props: FormProps): JSX.Element => {
    return (
        <form
            style={{ display: 'flex', justifyContent: 'space-between' }}
            method={props.form.submitMethod}
            action={props.form.submitAction}>
            {
                props.form.csrfToken !== undefined && props.form.csrfToken !== '' &&
                    <input type="hidden" name="csrf_token" value={props.form.csrfToken} />
            }
            { props.children }
        </form>
    );
};
