package rail

import "testing"

func TestManifestHashIgnoresDeclarationOrder(t *testing.T) {
	manifest := testManifest()
	reordered := testManifest()
	reordered.Components = []Component{manifest.Components[1], manifest.Components[0]}

	first, err := ManifestHash(manifest)
	if err != nil {
		t.Fatal(err)
	}
	second, err := ManifestHash(reordered)
	if err != nil {
		t.Fatal(err)
	}
	if first != second {
		t.Fatalf("reordering components changed the manifest hash: %s versus %s", first, second)
	}
}
