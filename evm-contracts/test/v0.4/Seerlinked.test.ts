import { matchers } from '@seerlink/test-helpers'
import { Seerlinked__factory } from '../../ethers/v0.4/factories/Seerlinked__factory'

const seerlinkedFactory = new Seerlinked__factory()

describe('Seerlinked', () => {
  it('has a limited public interface', async () => {
    matchers.publicAbi(seerlinkedFactory, [])
  })
})
