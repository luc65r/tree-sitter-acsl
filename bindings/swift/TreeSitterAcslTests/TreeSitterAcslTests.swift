import XCTest
import SwiftTreeSitter
import TreeSitterAcsl

final class TreeSitterAcslTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_acsl())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading ANSI/ISO C Specification Language grammar")
    }
}
