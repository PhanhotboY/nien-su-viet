import { redirect } from 'next/navigation';

const CmsdeskPage = async () => {
  redirect('/cmsdesk/historical-events');
};

export default CmsdeskPage;
